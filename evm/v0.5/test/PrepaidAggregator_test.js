import * as h from './support/helpers'
import { assertBigNum } from './support/matchers'

import { expectRevert } from 'openzeppelin-test-helpers'

contract('PrepaidAggregator', () => {
  const Aggregator = artifacts.require('PrepaidAggregator.sol')
  const personas = h.personas
  const paymentAmount = h.toWei('3')
  const deposit = h.toWei('100')
  const answer = 100
  const minAns = 1
  const maxAns = 1
  const rrDelay = 0

  let aggregator, link, nextRound, oracles

  function parseRound(log) {
    return {
      paymentAmount: h.bigNum(log.topics[1]),
      minAnswerCount: h.bigNum(log.topics[2]),
      maxAnswerCount: h.bigNum(log.topics[3]),
      restartDelay: h.bigNum(log.data),
    }
  }

  beforeEach(async () => {
    link = await h.linkContract(personas.defaultAccount)
    aggregator = await Aggregator.new(link.address, paymentAmount, {
      from: personas.Carol,
    })
    await link.transfer(aggregator.address, deposit)
    await aggregator.updateAvailableFunds()
    assertBigNum(deposit, await link.balanceOf.call(aggregator.address))
    nextRound = 1
  })

  it('has a limited public interface', () => {
    h.checkPublicABI(Aggregator, [
      'addOracle',
      'allocatedFunds',
      'availableFunds',
      'currentAnswer',
      'currentRound',
      'getAnswer',
      'getUpdatedHeight',
      'latestRound',
      'latestSubmission',
      'maxAnswerCount',
      'minAnswerCount',
      'oracleCount',
      'paymentAmount',
      'removeOracle',
      'restartDelay',
      'updateAnswer',
      'updateAvailableFunds',
      'updateFutureRounds',
      'updatedHeight',
      'withdraw',
      'withdrawFunds',
      'withdrawable',
      // Ownable methods:
      'isOwner',
      'owner',
      'transferOwnership',
    ])
  })

  describe('#updateAnswer', async () => {
    let minMax

    beforeEach(async () => {
      oracles = [personas.Neil, personas.Ned, personas.Nelly]
      for (let i = 0; i < oracles.length; i++) {
        minMax = i + 1
        await aggregator.addOracle(oracles[i], minMax, minMax, rrDelay, {
          from: personas.Carol,
        })
      }
    })

    it('updates the allocated and available funds counters', async () => {
      assertBigNum(0, await aggregator.allocatedFunds.call())

      const tx = await aggregator.updateAnswer(nextRound, answer, {
        from: personas.Neil,
      })

      assertBigNum(paymentAmount, await aggregator.allocatedFunds.call())
      const expectedAvailable = deposit.sub(paymentAmount)
      assertBigNum(expectedAvailable, await aggregator.availableFunds.call())
      const logged = h.bigNum(tx.receipt.rawLogs[1].topics[1])
      assertBigNum(expectedAvailable, logged)
    })

    it('updates the latest submission record for the oracle', async () => {
      let latest = await aggregator.latestSubmission.call(personas.Neil)
      assert.equal(0, latest[0])
      assert.equal(0, latest[1])

      const newAnswer = 427
      await aggregator.updateAnswer(nextRound, newAnswer, {
        from: personas.Neil,
      })

      latest = await aggregator.latestSubmission.call(personas.Neil)
      assert.equal(newAnswer, latest[0])
      assert.notEqual(0, latest[1])
    })

    context('when the minimum oracles have not reported', async () => {
      it('pays the oracles that have reported', async () => {
        assertBigNum(
          0,
          await aggregator.withdrawable.call({ from: personas.Neil }),
        )

        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Neil,
        })

        assertBigNum(
          paymentAmount,
          await aggregator.withdrawable.call({ from: personas.Neil }),
        )
        assertBigNum(
          0,
          await aggregator.withdrawable.call({ from: personas.Ned }),
        )
        assertBigNum(
          0,
          await aggregator.withdrawable.call({ from: personas.Nelly }),
        )
      })

      it('does not update the answer', async () => {
        assert.equal(0, await aggregator.currentAnswer.call())

        // Not updated because of changes by the owner setting minAnswerCount to 3
        await aggregator.updateAnswer(nextRound, answer, { from: personas.Ned })
        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Nelly,
        })

        assert.equal(0, await aggregator.currentAnswer.call())
      })
    })

    context('when the minimum number of oracles have reported', async () => {
      beforeEach(async () => {
        await aggregator.updateFutureRounds(paymentAmount, 2, 3, 0, {
          from: personas.Carol,
        })
        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Neil,
        })
      })

      it('updates the answer with the median', async () => {
        assert.equal(0, await aggregator.currentAnswer.call())

        await aggregator.updateAnswer(nextRound, 99, { from: personas.Ned })
        assert.equal(99, await aggregator.currentAnswer.call()) // ((100+99) / 2).to_i

        await aggregator.updateAnswer(nextRound, 101, {
          from: personas.Nelly,
        })

        assert.equal(100, await aggregator.currentAnswer.call())
      })

      it('updates the updated height', async () => {
        const originalHeight = await aggregator.updatedHeight.call()
        assert.equal(0, originalHeight.toNumber())

        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Nelly,
        })

        const currentHeight = await aggregator.updatedHeight.call()
        assert.isAbove(currentHeight.toNumber(), originalHeight.toNumber())
      })

      it('announces the new answer with a log event', async () => {
        const tx = await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Nelly,
        })
        const log = tx.receipt.rawLogs[0]
        const newAnswer = h.bigNum(log.topics[1])

        assert.equal(answer, newAnswer.toNumber())
      })
    })

    context('when an oracle submits for a round twice', async () => {
      it('reverts', async () => {
        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Neil,
        })

        await expectRevert(
          aggregator.updateAnswer(nextRound, answer, {
            from: personas.Neil,
          }),
          'Cannot update round reports',
        )
      })
    })

    context('when updated after the max answers submitted', async () => {
      beforeEach(async () => {
        await aggregator.updateFutureRounds(paymentAmount, 1, 1, rrDelay, {
          from: personas.Carol,
        })
        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Neil,
        })
      })

      it('reverts', async () => {
        await expectRevert(
          aggregator.updateAnswer(nextRound, answer, {
            from: personas.Ned,
          }),
          'Max responses reached for round',
        )
      })
    })

    context('when a new highest round number is passed in', async () => {
      it('increments the answer round', async () => {
        assert.equal(0, await aggregator.currentRound.call())

        for (const oracle of oracles) {
          await aggregator.updateAnswer(nextRound, answer, { from: oracle })
        }

        assert.equal(1, await aggregator.currentRound.call())
      })

      it('announces a new round by emitting a log', async () => {
        const tx = await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Neil,
        })
        const log = tx.receipt.rawLogs[0]
        const roundNumber = h.bigNum(log.topics[1])
        const startedBy = h.toAddress(log.topics[2])

        assert.equal(nextRound, roundNumber.toNumber())
        assert.equal(startedBy, personas.Neil)
      })
    })

    context('when a round is passed in higher than expected', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.updateAnswer(nextRound + 1, answer, {
            from: personas.Neil,
          }),
          'Must report on current round',
        )
      })
    })

    context('when called by a non-oracle', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.updateAnswer(nextRound, answer, {
            from: personas.Carol,
          }),
          'Only updatable by designated oracles',
        )
      })
    })

    context('when there are not sufficient available funds', async () => {
      beforeEach(async () => {
        await aggregator.withdrawFunds(personas.Carol, deposit, {
          from: personas.Carol,
        })
      })

      it('reverts', async () => {
        await expectRevert(
          aggregator.updateAnswer(nextRound, answer, {
            from: personas.Neil,
          }),
          'SafeMath: subtraction overflow',
        )
      })
    })

    context('when price is updated mid-round', async () => {
      const newAmount = h.toWei('50')

      it('pays the same amount to all oracles per round', async () => {
        assertBigNum(
          0,
          await aggregator.withdrawable.call({ from: personas.Neil }),
        )
        assertBigNum(
          0,
          await aggregator.withdrawable.call({ from: personas.Nelly }),
        )

        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Neil,
        })

        await aggregator.updateFutureRounds(
          newAmount,
          minMax,
          minMax,
          rrDelay,
          { from: personas.Carol },
        )

        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Nelly,
        })

        assertBigNum(
          paymentAmount,
          await aggregator.withdrawable.call({ from: personas.Neil }),
        )
        assertBigNum(
          paymentAmount,
          await aggregator.withdrawable.call({ from: personas.Nelly }),
        )
      })
    })

    context(
      'when an oracle starts a round before the restart delay is over',
      async () => {
        beforeEach(async () => {
          await aggregator.updateFutureRounds(paymentAmount, 1, 1, 0, {
            from: personas.Carol,
          })

          oracles = [personas.Neil, personas.Ned, personas.Nelly]
          for (let i = 0; i < oracles.length; i++) {
            await aggregator.updateAnswer(nextRound, answer, {
              from: oracles[i],
            })
            nextRound++
          }

          const newDelay = 2
          // Since Ned and Nelly have answered recently, and we set the delay
          // to 2, only Nelly can answer as she is the only oracle that hasn't
          // started the last two rounds.
          await aggregator.updateFutureRounds(
            paymentAmount,
            1,
            oracles.length,
            newDelay,
            {
              from: personas.Carol,
            },
          )
        })

        context(
          'when called by an oracle who has not answered recently',
          async () => {
            it('does not revert', async () => {
              await aggregator.updateAnswer(nextRound, answer, {
                from: personas.Neil,
              })
            })
          },
        )

        context('when called by an oracle who answered recently', async () => {
          it('reverts', async () => {
            await expectRevert(
              aggregator.updateAnswer(nextRound, answer, {
                from: personas.Ned,
              }),
              'Max responses reached for round',
            )

            await expectRevert(
              aggregator.updateAnswer(nextRound, answer, {
                from: personas.Nelly,
              }),
              'Max responses reached for round',
            )
          })
        })
      },
    )
  })

  describe('#getAnswer', async () => {
    const answers = [1, 10, 101, 1010, 10101, 101010, 1010101]

    beforeEach(async () => {
      await aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
        from: personas.Carol,
      })

      for (const answer of answers) {
        await aggregator.updateAnswer(nextRound, answer, {
          from: personas.Neil,
        })
        nextRound++
      }
    })

    it('retrieves the answer recorded for past rounds', async () => {
      for (let i = nextRound; i < nextRound; i++) {
        const answer = await aggregator.getAnswer.call(i)
        assertBigNum(h.bigNum(answers[i - 1]), answer)
      }
    })
  })

  describe('#getUpdatedHeight', async () => {
    beforeEach(async () => {
      await aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
        from: personas.Carol,
      })

      for (let i = 0; i < 10; i++) {
        await aggregator.updateAnswer(nextRound, i, {
          from: personas.Neil,
        })
        nextRound++
      }
    })

    it('retrieves the answer recorded for past rounds', async () => {
      let lastHeight = h.bigNum(0)

      for (let i = 1; i < nextRound; i++) {
        const currentHeight = await aggregator.getUpdatedHeight.call(i)
        assert.isAbove(currentHeight.toNumber(), lastHeight.toNumber())
        lastHeight = currentHeight
      }
    })
  })

  describe('#addOracle', async () => {
    it('increases the oracle count', async () => {
      const pastCount = await aggregator.oracleCount.call()
      await aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
        from: personas.Carol,
      })
      const currentCount = await aggregator.oracleCount.call()

      assertBigNum(currentCount, pastCount.add(h.bigNum(1)))
    })

    it('updates the round details', async () => {
      await aggregator.addOracle(personas.Neil, 0, 1, 0, {
        from: personas.Carol,
      })

      assertBigNum(h.bigNum(0), await aggregator.minAnswerCount.call())
      assertBigNum(h.bigNum(1), await aggregator.maxAnswerCount.call())
      assertBigNum(h.bigNum(0), await aggregator.restartDelay.call())
    })

    it('emits a log', async () => {
      const tx = await aggregator.addOracle(
        personas.Neil,
        minAns,
        maxAns,
        rrDelay,
        {
          from: personas.Carol,
        },
      )

      const added = h.toAddress(tx.receipt.rawLogs[0].topics[1])
      assertBigNum(added, personas.Neil)
    })

    context('when the oracle has already been added', async () => {
      beforeEach(async () => {
        await aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
          from: personas.Carol,
        })
      })

      it('reverts', async () => {
        await expectRevert(
          aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
            from: personas.Carol,
          }),
          'Address is already recorded as an oracle',
        )
      })
    })

    context('when called by anyone but the owner', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
            from: personas.Neil,
          }),
          'Ownable: caller is not the owner',
        )
      })
    })

    const limit = 42
    context(`when adding more than ${limit} oracles`, async () => {
      it('reverts', async () => {
        for (let i = 0; i < limit; i++) {
          const minMax = i + 1
          const fakeAddress = web3.utils.randomHex(20)
          await aggregator.addOracle(fakeAddress, minMax, minMax, rrDelay, {
            from: personas.Carol,
          })
        }
        await expectRevert(
          aggregator.addOracle(personas.Neil, limit + 1, limit + 1, rrDelay, {
            from: personas.Carol,
          }),
          `cannot add more than ${limit} oracles`,
        )
      })
    })
  })

  describe('#removeOracle', async () => {
    beforeEach(async () => {
      await aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
        from: personas.Carol,
      })
      await aggregator.addOracle(personas.Nelly, 2, 2, rrDelay, {
        from: personas.Carol,
      })
    })

    it('decreases the oracle count', async () => {
      const pastCount = await aggregator.oracleCount.call()
      await aggregator.removeOracle(personas.Neil, minAns, maxAns, rrDelay, {
        from: personas.Carol,
      })
      const currentCount = await aggregator.oracleCount.call()

      assertBigNum(currentCount, pastCount.sub(h.bigNum(1)))
    })

    it('updates the round details', async () => {
      await aggregator.removeOracle(personas.Neil, 0, 1, 0, {
        from: personas.Carol,
      })

      assertBigNum(h.bigNum(0), await aggregator.minAnswerCount.call())
      assertBigNum(h.bigNum(1), await aggregator.maxAnswerCount.call())
      assertBigNum(h.bigNum(0), await aggregator.restartDelay.call())
    })

    it('emits a log', async () => {
      const tx = await aggregator.removeOracle(
        personas.Neil,
        minAns,
        maxAns,
        rrDelay,
        {
          from: personas.Carol,
        },
      )

      const added = h.toAddress(tx.receipt.rawLogs[0].topics[1])
      assertBigNum(added, personas.Neil)
    })

    context('when the oracle is not currently added', async () => {
      beforeEach(async () => {
        await aggregator.removeOracle(personas.Neil, minAns, maxAns, rrDelay, {
          from: personas.Carol,
        })
      })

      it('reverts', async () => {
        await expectRevert(
          aggregator.removeOracle(personas.Neil, minAns, maxAns, rrDelay, {
            from: personas.Carol,
          }),
          'Address is not an oracle',
        )
      })
    })

    context('when removing the last oracle', async () => {
      it('does not revert', async () => {
        await aggregator.removeOracle(personas.Neil, minAns, maxAns, rrDelay, {
          from: personas.Carol,
        })

        await aggregator.removeOracle(personas.Nelly, 0, 0, 0, {
          from: personas.Carol,
        })
      })
    })

    context('when called by anyone but the owner', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.removeOracle(personas.Neil, 0, 0, rrDelay, {
            from: personas.Ned,
          }),
          'Ownable: caller is not the owner',
        )
      })
    })
  })

  describe('#withdrawFunds', () => {
    context('when called by the owner', () => {
      it('succeeds', async () => {
        await aggregator.withdrawFunds(personas.Carol, deposit, {
          from: personas.Carol,
        })

        assertBigNum(0, await aggregator.availableFunds.call())
        assertBigNum(deposit, await link.balanceOf.call(personas.Carol))
      })

      it('does not let withdrawals happen multiple times', async () => {
        await aggregator.withdrawFunds(personas.Carol, deposit, {
          from: personas.Carol,
        })

        await expectRevert(
          aggregator.withdrawFunds(personas.Carol, deposit, {
            from: personas.Carol,
          }),
          'Insufficient funds',
        )
      })

      context('with a number higher than the available LINK balance', () => {
        beforeEach(async () => {
          await aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
            from: personas.Carol,
          })
          await aggregator.updateAnswer(nextRound, answer, {
            from: personas.Neil,
          })
        })

        it('fails', async () => {
          await expectRevert(
            aggregator.withdrawFunds(personas.Carol, deposit, {
              from: personas.Carol,
            }),
            'Insufficient funds',
          )

          assertBigNum(
            deposit.sub(paymentAmount),
            await aggregator.availableFunds.call(),
          )
        })
      })
    })

    context('when called by a non-owner', () => {
      it('fails', async () => {
        await expectRevert(
          aggregator.withdrawFunds(personas.Carol, deposit, {
            from: personas.Eddy,
          }),
          'Ownable: caller is not the owner',
        )

        assertBigNum(deposit, await aggregator.availableFunds.call())
      })
    })
  })

  describe('#updateFutureRounds', async () => {
    let minAnswerCount, maxAnswerCount
    const newPaymentAmount = h.toWei('2')
    const newMin = 1
    const newMax = 3
    const newDelay = 2

    beforeEach(async () => {
      oracles = [personas.Neil, personas.Ned, personas.Nelly]
      for (let i = 0; i < oracles.length; i++) {
        const minMax = i + 1
        await aggregator.addOracle(oracles[i], minMax, minMax, rrDelay, {
          from: personas.Carol,
        })
      }
      minAnswerCount = oracles.length
      maxAnswerCount = oracles.length

      assertBigNum(paymentAmount, await aggregator.paymentAmount.call())
      assert.equal(minAnswerCount, await aggregator.minAnswerCount.call())
      assert.equal(maxAnswerCount, await aggregator.maxAnswerCount.call())
    })

    it('updates the min and max answer counts', async () => {
      await aggregator.updateFutureRounds(
        newPaymentAmount,
        newMin,
        newMax,
        newDelay,
        {
          from: personas.Carol,
        },
      )

      assertBigNum(newPaymentAmount, await aggregator.paymentAmount.call())
      assertBigNum(h.bigNum(newMin), await aggregator.minAnswerCount.call())
      assertBigNum(h.bigNum(newMax), await aggregator.maxAnswerCount.call())
      assertBigNum(h.bigNum(newDelay), await aggregator.restartDelay.call())
    })

    it('emits a log announcing the new round details', async () => {
      const tx = await aggregator.updateFutureRounds(
        paymentAmount,
        newMin,
        newMax,
        newDelay,
        {
          from: personas.Carol,
        },
      )

      const round = parseRound(tx.receipt.rawLogs[0])

      assertBigNum(paymentAmount, round.paymentAmount)
      assertBigNum(h.bigNum(newMin), round.minAnswerCount)
      assertBigNum(h.bigNum(newMax), round.maxAnswerCount)
      assertBigNum(h.bigNum(newDelay), round.restartDelay)
    })

    context('when it is set to higher than the number or oracles', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.updateFutureRounds(
            paymentAmount,
            minAnswerCount,
            4,
            rrDelay,
            {
              from: personas.Carol,
            },
          ),
          'Cannot have the answer max higher oracle count',
        )
      })
    })

    context('when it sets the min higher than the max', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.updateFutureRounds(paymentAmount, 3, 2, rrDelay, {
            from: personas.Carol,
          }),
          'Cannot have the answer minimum higher the max',
        )
      })
    })

    context('when delay equal or greater the oracle count', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.updateFutureRounds(paymentAmount, 1, 1, 3, {
            from: personas.Carol,
          }),
          'Restart delay must be less than oracle count',
        )
      })
    })

    context('when called by anyone but the owner', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.updateFutureRounds(paymentAmount, 1, 3, rrDelay, {
            from: personas.Ned,
          }),
          'caller is not the owner',
        )
      })
    })
  })

  describe('#updateAvailableFunds', async () => {
    it('checks the LINK token to see if any additional funds are available', async () => {
      const originalBalance = await aggregator.availableFunds.call()

      await aggregator.updateAvailableFunds()

      assertBigNum(originalBalance, await aggregator.availableFunds.call())

      await link.transfer(aggregator.address, deposit)
      await aggregator.updateAvailableFunds()

      const newBalance = await aggregator.availableFunds.call()
      assertBigNum(originalBalance.add(deposit), newBalance)
    })

    it('removes allocated funds from the available balance', async () => {
      const originalBalance = await aggregator.availableFunds.call()

      await aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
        from: personas.Carol,
      })
      await aggregator.updateAnswer(nextRound, answer, {
        from: personas.Neil,
      })
      await link.transfer(aggregator.address, deposit)
      await aggregator.updateAvailableFunds()

      const expected = originalBalance.add(deposit).sub(paymentAmount)
      const newBalance = await aggregator.availableFunds.call()
      assertBigNum(expected, newBalance)
    })

    it('emits a log', async () => {
      await link.transfer(aggregator.address, deposit)

      const tx = await aggregator.updateAvailableFunds()

      const reportedBalance = h.bigNum(tx.receipt.rawLogs[0].topics[1])
      assertBigNum(await aggregator.availableFunds.call(), reportedBalance)
    })
  })

  describe('#withdraw', async () => {
    beforeEach(async () => {
      await aggregator.addOracle(personas.Neil, minAns, maxAns, rrDelay, {
        from: personas.Carol,
      })
      await aggregator.updateAnswer(nextRound, answer, {
        from: personas.Neil,
      })
    })

    it('transfers LINK to the caller', async () => {
      const originalBalance = await link.balanceOf.call(aggregator.address)
      assertBigNum(0, await link.balanceOf.call(personas.Neil))

      await aggregator.withdraw(personas.Neil, paymentAmount, {
        from: personas.Neil,
      })

      assertBigNum(
        originalBalance.sub(paymentAmount),
        await link.balanceOf.call(aggregator.address),
      )
      assertBigNum(paymentAmount, await link.balanceOf.call(personas.Neil))
    })

    it('decrements the allocated funds counter', async () => {
      const originalAllocation = await aggregator.allocatedFunds.call()

      await aggregator.withdraw(personas.Neil, paymentAmount, {
        from: personas.Neil,
      })

      assertBigNum(
        originalAllocation.sub(paymentAmount),
        await aggregator.allocatedFunds.call(),
      )
    })

    context('when the caller withdraws more than they have', async () => {
      it('reverts', async () => {
        await expectRevert(
          aggregator.withdraw(personas.Neil, paymentAmount.add(h.bigNum(1)), {
            from: personas.Neil,
          }),
          'Insufficient balance',
        )
      })
    })
  })
})
