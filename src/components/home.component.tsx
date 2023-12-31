'use client'

import { TNewApplication } from '@/app/types'
import { createPool } from '@/sdk/allo'
import { allocate, createApplication } from '@/sdk/microgrants'
import { createProfile } from '@/sdk/registry'
import { chainData, wagmiConfigData } from '@/services/wagmi'
import { Allocation } from '@allo-team/allo-v2-sdk/dist/strategies/MicroGrantsStrategy/types'
import { Status } from '@allo-team/allo-v2-sdk/dist/strategies/types'
import { ConnectButton, RainbowKitProvider, midnightTheme } from '@rainbow-me/rainbowkit'
import '@rainbow-me/rainbowkit/styles.css'
import { WagmiConfig } from 'wagmi'

export function Home() {
  // Set this here so we dont have to create a new profile every time and we are not managing state in this demo.
  // We use the profileId to create a new application in `_newApplicationData`.
  const profileId = '0xbbc8d31c2b00a011912322740d238139ae8578a4dec18818408f720bdcb54b53'
  const _newApplicationData: TNewApplication = {
    name: 'MunayFund App',
    website: 'https://munayfund.vercel.app',
    description: 'MunayFund Description',
    email: 'wolfcito.eth+munayfund@gmail.com',
    requestedAmount: BigInt(1e12),
    recipientAddress: '0x98aCac2dc4bb94748e33f092D7D70911b38AB76b',
    base64Image: '',
    profileName: 'Wolfcito',
    profileId: profileId,
  }

  const _allocationData: Allocation = {
    recipientId: '0x98aCac2dc4bb94748e33f092D7D70911b38AB76b',
    status: Status.Accepted,
  }

  return (
    <WagmiConfig config={wagmiConfigData}>
      <RainbowKitProvider chains={chainData} modalSize="wide" theme={midnightTheme()}>
        <main className="flex min-h-screen flex-col items-center justify-between p-2">
          <div className="flex w-full max-w-5xl items-center justify-between font-mono text-sm">
            Munayfund
            <ConnectButton />
          </div>

          <div className="flex flex-row">
            <button
              onClick={() =>
                createProfile().then((res: any) => {
                  console.log('Profile ID: ', res)
                  alert('Profile created with ID: ' + res)
                })
              }
              className="mx-2 rounded-lg bg-gradient-to-r from-[#ff00a0] to-[#d75fab] px-4 py-2 text-white"
            >
              Create Profile
            </button>
            <button
              onClick={() => {
                createPool().then((res: any) => {
                  console.log('Pool ID: ', res.poolId)
                  alert('Pool created with ID: ' + res.poolId)
                })
              }}
              className="mx-2 rounded-lg bg-gradient-to-r from-[#ff00a0] to-[#d75fab] px-4 py-2 text-white"
            >
              Create Pool
            </button>
            <button
              onClick={() => {
                createApplication(_newApplicationData, 421614, 12).then((res: any) => {
                  console.log('Recipient ID: ', res.recipientId)
                  alert('Applied with ID: ' + res.recipientId)
                })
              }}
              className="mx-2 rounded-lg bg-gradient-to-r from-[#ff00a0] to-[#d75fab] px-4 py-2 text-white"
            >
              Apply to Pool
            </button>
            {/* WIP */}
            <button
              onClick={() => {
                allocate(_allocationData).then((res: any) => {
                  console.log('Recipient ID: ', res.recipientId)
                  alert('Applied with ID: ' + res.recipientId)
                })
              }}
              className="mx-2 rounded-lg bg-gradient-to-r from-[#ff00a0] to-[#d75fab] px-4 py-2 text-white"
            >
              Allocate to Pool
            </button>
          </div>

          <div className="-mt-40">Powered by @Kanicrafters</div>
        </main>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}
