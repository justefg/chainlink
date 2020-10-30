import React from 'react'
import { Route } from 'react-router-dom'
import { jsonApiOcrKeys, OcrKeyBundle } from 'factories/jsonApiOcrKeys'
import { syncFetch } from 'test-helpers/syncFetch'
import globPath from 'test-helpers/globPath'
import { mountWithProviders } from 'test-helpers/mountWithTheme'
import { partialAsFull } from '@chainlink/ts-helpers'

import {
  ENDPOINT as OCR_ENDPOINT,
  INDEX_ENDPOINT as OCR_INDEX_ENDPOINT,
} from 'api/v2/ocrKeys'
import { KeysIndex } from 'pages/Keys/Index'

describe('pages/Keys/Index', () => {
  describe('Off-Chain Reporting keys', () => {
    it('renders the list of keys', async () => {
      const [expectedKey1, expectedKey2] = [
        partialAsFull<OcrKeyBundle>({
          id: 'keyId1',
          createdAt: new Date().toISOString(),
          offChainPublicKey: 'offChainPublicKey1',
          configPublicKey: 'configPublicKey1',
          onChainSigningAddress: 'onChainSigningAddress1',
        }),
        partialAsFull<OcrKeyBundle>({
          id: 'keyId2',
          createdAt: new Date().toISOString(),
          offChainPublicKey: 'offChainPublicKey2',
          configPublicKey: 'configPublicKey2',
          onChainSigningAddress: 'onChainSigningAddress2',
        }),
      ]

      global.fetch.getOnce(
        globPath(OCR_INDEX_ENDPOINT),
        jsonApiOcrKeys([expectedKey1, expectedKey2]),
      )

      const wrapper = mountWithProviders(<Route component={KeysIndex} />)
      await syncFetch(wrapper)
      expect(wrapper.text()).toContain('just now')
      expect(wrapper.text()).toContain('Delete')
      expect(wrapper.find('tbody').children().length).toEqual(2)
      expect(wrapper.text()).toContain(expectedKey1.id)
      expect(wrapper.text()).toContain(expectedKey1.offChainPublicKey)
      expect(wrapper.text()).toContain(expectedKey1.configPublicKey)
      expect(wrapper.text()).toContain(expectedKey1.onChainSigningAddress)
      expect(wrapper.text()).toContain(expectedKey1.offChainPublicKey)
      expect(wrapper.text()).toContain(expectedKey2.id)
      expect(wrapper.text()).toContain(expectedKey2.offChainPublicKey)
      expect(wrapper.text()).toContain(expectedKey2.configPublicKey)
      expect(wrapper.text()).toContain(expectedKey2.onChainSigningAddress)
      expect(wrapper.text()).toContain(expectedKey2.offChainPublicKey)
    })

    it('allows to create a new key bundle', async () => {
      const expectedKey = partialAsFull<OcrKeyBundle>({
        id: 'keyId',
      })
      global.fetch.getOnce(globPath(OCR_INDEX_ENDPOINT), [])
      const wrapper = mountWithProviders(<Route component={KeysIndex} />)
      await syncFetch(wrapper)

      expect(wrapper.find('tbody').children().length).toEqual(0)
      expect(wrapper.text()).not.toContain(expectedKey.id)

      global.fetch.getOnce(
        globPath(OCR_INDEX_ENDPOINT),
        jsonApiOcrKeys([expectedKey]),
      )
      global.fetch.postOnce(globPath(OCR_ENDPOINT), {})
      wrapper.find('[data-testid="keys-ocr-create"]').first().simulate('click')
      await syncFetch(wrapper)

      expect(wrapper.find('tbody').children().length).toEqual(1)
      expect(wrapper.text()).toContain(expectedKey.id)
    })

    it('allows to delete a key bundle', async () => {
      const expectedKey = partialAsFull<OcrKeyBundle>({
        id: 'keyId',
      })
      global.fetch.getOnce(
        globPath(OCR_INDEX_ENDPOINT),
        jsonApiOcrKeys([expectedKey]),
      )
      const wrapper = mountWithProviders(<Route component={KeysIndex} />)
      await syncFetch(wrapper)

      expect(wrapper.find('tbody').children().length).toEqual(1)
      expect(wrapper.text()).toContain(expectedKey.id)

      global.fetch.getOnce(globPath(OCR_INDEX_ENDPOINT), {})
      global.fetch.deleteOnce(
        globPath(`${OCR_INDEX_ENDPOINT}/${expectedKey.id}`),
        {},
      )
      wrapper
        .find('[data-testid="keys-ocr-delete-dialog"]')
        .first()
        .simulate('click')
      wrapper
        .find('[data-testid="keys-ocr-delete-confirm"]')
        .first()
        .simulate('click')
      await syncFetch(wrapper)

      expect(wrapper.find('tbody').children().length).toEqual(0)
      expect(wrapper.text()).not.toContain(expectedKey.id)
    })
  })
})
