// package main is a script for generating geth golang contract wrappers for
// solidity artifacts generated by belt.
//
//  Usage:
//
//  go run wrap.go <sol-compiler-output-path> <package-name>
//
// This will output the generated file to
// ../generated/<package-name>/<package-name>.go

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers"

	gethParams "github.com/ethereum/go-ethereum/params"
)

func main() {
	beltArtifactPath := os.Args[1] // path to belt artifact to wrap
	pkgName := os.Args[2]          // golang package name for wrapper

	contract, err := gethwrappers.ExtractContractDetails(beltArtifactPath)
	if err != nil {
		exit("could not get contract details from belt artifact", err)
	}
	// It might seem like a shame, to shell out to another golang program like
	// this, but the abigen executable is the stable public interface to the
	// geth contract-wrapper machinery.
	//
	// Check whether native abigen is installed, and has correct version
	var versionResponse bytes.Buffer
	abigenExecutablePath := filepath.Join(getProjectRoot(), "tools/bin/abigen")
	abigenVersionCheck := exec.Command(abigenExecutablePath, "--version")
	abigenVersionCheck.Stdout = &versionResponse
	if err := abigenVersionCheck.Run(); err != nil {
		exit("no native abigen; you must install it (`make abigen` in the "+
			"chainlink root dir)", err)
	}
	version := string(regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+`).Find(
		versionResponse.Bytes()))
	if version != gethParams.Version {
		exit(fmt.Sprintf("wrong version (%s) of abigen; install the correct one "+
			"(%s) with `make abigen` in the chainlink root dir", version,
			gethParams.Version),
			nil)
	}

	className := filepath.Base(beltArtifactPath)
	if !strings.HasSuffix(className, ".json") {
		exit("belt artifact path should end with `.json`: "+className, nil)
	}
	className = className[:len(className)-len(".json")]
	tmpDir, err := ioutil.TempDir("", className+"-contractWrapper")
	if err != nil {
		exit("failed to create temporary working directory", err)
	}
	defer func(tmpDir string) {
		if err := os.RemoveAll(tmpDir); err != nil {
			fmt.Println("failure while cleaning up temporary working directory:", err)
		}
	}(tmpDir)

	binPath := filepath.Join(tmpDir, "bin")
	if err := ioutil.WriteFile(binPath, []byte(contract.Binary), 0600); err != nil {
		exit("could not write contract binary to temp working directory", err)
	}
	abiPath := filepath.Join(tmpDir, "abi")
	if err := ioutil.WriteFile(abiPath, []byte(contract.ABI), 0600); err != nil {
		exit("could not write contract binary to temp working directory", err)
	}
	cwd, err := os.Getwd() // gethwrappers directory
	if err != nil {
		exit("could not get working directory", err)
	}
	outPath := filepath.Join(cwd, "generated", pkgName, pkgName+".go")

	buildCommand := exec.Command(
		abigenExecutablePath,
		"-bin", binPath,
		"-abi", abiPath,
		"-out", outPath,
		"-type", className,
		"-pkg", pkgName,
	)
	var buildResponse bytes.Buffer
	buildCommand.Stderr = &buildResponse
	if err := buildCommand.Run(); err != nil {
		exit("failure while building "+className+" wrapper, stderr: "+
			string(buildResponse.Bytes()), err)
	}

	// Build succeeded, so update the versions db with the new contract data
	versions, err := gethwrappers.ReadVersionsDB()
	if err != nil {
		exit("could not read current versions database", err)
	}
	versions.GethVersion = gethParams.Version
	versions.ContractVersions[pkgName] = gethwrappers.ContractVersion{
		CompilerArtifactPath: beltArtifactPath,
		Hash:                 contract.VersionHash(),
	}
	if err := gethwrappers.WriteVersionsDB(versions); err != nil {
		exit("could not save versions db", err)
	}
}

func exit(msg string, err error) {
	if err != nil {
		fmt.Println(msg+":", err)
	} else {
		fmt.Println(msg)
	}
	os.Exit(1)
}

// getProjectRoot returns the root of the chainlink project
func getProjectRoot() (rootPath string) {
	root, err := os.Getwd()
	if err != nil {
		exit("could not get current working directory while seeking project root",
			err)
	}
	for root != "/" { // Walk up path to find dir containing go.mod
		if _, err := os.Stat(filepath.Join(root, "go.mod")); os.IsNotExist(err) {
			root = filepath.Dir(root)
		} else {
			return root
		}
	}
	exit("could not find project root", nil)
	panic("can't get here")
}
