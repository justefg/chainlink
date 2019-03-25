import 'reflect-metadata'
import { createConnection, Connection } from 'typeorm'
import { PostgresConnectionOptions } from 'typeorm/driver/postgres/PostgresConnectionOptions'
import options from '../ormconfig.json'

const overridableKeys = ['host', 'port', 'username', 'password', 'database']

const isEnvEqual = (optionName: string, env: string): boolean => {
  return env == optionName ||
    (env == "development" && optionName == "default")
}

const loadOptions = (env = process.env.NODE_ENV || "development") => {
  for (let option of options) {
    if (isEnvEqual(option.name, env)) {
      return option
    }
  }
  throw Error(`env ${env} not found in options from ormconfig.json`)
}

// Loads the following ENV vars, giving them precedence.
// i.e. TYPEORM_PORT will replace "port" in ormconfig.json.
const mergeOptions = (): PostgresConnectionOptions => {
  let envOptions: { [key: string]: string } = {}
  for (let v of overridableKeys) {
    const envVar = process.env[`TYPEORM_${v.toUpperCase()}`]
    if (envVar) {
      envOptions[v] = envVar
    }
  }
  return {
    ...loadOptions(),
    ...envOptions
  } as PostgresConnectionOptions
}

let db: Connection
export const createDbConnection = async (): Promise<Connection> => {
  db = await createConnection(mergeOptions())
  return db
}

export const getDb = (): Connection => {
  return db
}
