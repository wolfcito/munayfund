'use client'

import { TNewApplication } from '@/app/types'
import { createPool } from '@/sdk/allo'
import { createApplication } from '@/sdk/microgrants'
import { createProfile } from '@/sdk/registry'
import { chainData, wagmiConfigData } from '@/services/wagmi'
import { Allocation } from '@allo-team/allo-v2-sdk/dist/strategies/MicroGrantsStrategy/types'
import { Status } from '@allo-team/allo-v2-sdk/dist/strategies/types'
import { ConnectButton, RainbowKitProvider, midnightTheme } from '@rainbow-me/rainbowkit'
import '@rainbow-me/rainbowkit/styles.css'
import { WagmiConfig } from 'wagmi'
import { ButtonNeon, ButtonSimple } from '../button'
import Image from 'next/image'
import Link from 'next/link'
import { Header } from '../header'
import clsx from 'clsx'

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
      <RainbowKitProvider chains={chainData} modalSize="wide" theme={midnightTheme()} coolMode>
        <main className="flex flex-col my-2 mx-7 min-h-screen">
          <Header />

          <Article>
            <div className="flex flex-col max-w-2xl">
              <Title label="Transparent, Secure, and Fueled by Innovation" />

              <p className="font-merienda text-2xl py-10">Decentralizing Dreams, Empowering Innovations.</p>
              <p className="text-slate-300 py-10">
                Explore, Fund, Decide - Revolutionizing collective financing with Arbitrum and Allo Protocol. Join us in
                a seamless journey of creation, funding, and decision-making.
              </p>
              <p className="py-5 space-x-5">
                <ButtonNeon
                  onClick={() =>
                    createProfile().then((res: any) => {
                      console.log('Profile ID: ', res)
                      alert('Profile created with ID: ' + res)
                    })
                  }
                  label="Create Profile"
                />

                <ButtonNeon
                  onClick={() => {
                    createPool().then((res: any) => {
                      console.log('Pool ID: ', res.poolId)
                      alert('Pool created with ID: ' + res.poolId)
                    })
                  }}
                  label="Create Pool"
                />

                <ButtonNeon
                  onClick={() => {
                    createApplication(_newApplicationData, 421614, 12).then((res: any) => {
                      console.log('Recipient ID: ', res.recipientId)
                      alert('Applied with ID: ' + res.recipientId)
                    })
                  }}
                  label="Apply to Pool"
                />
              </p>
            </div>

            <Image src="/munayfund.png" alt="munayfund image" width={200} height={300} className="h-auto w-80" />
          </Article>

          <Article className="min-h-screen">
            <div className="flex flex-col max-w-lg space-y-4">
              <Title label="About Munayfund" />

              <p className="text-slate-300">
                Our crowdfunding platform, powered by Arbitrum and Allo Protocol, revolutionizes decentralized funding,
                simplifying project creation and funding on the Arbitrum network.
              </p>
              <p className="text-slate-300">
                With a common fund managed by Allo Protocol, users vote for fund distribution, and the platform
                integrates machine learning, an intuitive interface, and robust security for transparent community
                engagement in decentralized finance evolution.
              </p>
            </div>

            <Image src="/upcurve.jpg" alt="levelup" width={100} height={300} className="h-auto w-52 ml-20" />
            <div className="px-4 space-y-5">
              <div className="text-right">
                <h3 className="text-balance font-bold text-[3em] leading-tight text-limelight">+2M</h3>
                <p className="text-slate-300">Funding volume</p>
              </div>
              <div className="text-right">
                <h3 className="text-balance font-bold text-[3em] leading-tight text-limelight">+50</h3>
                <p className="text-slate-300">Completed Projects</p>
              </div>
              <div className="text-right">
                <h3 className="text-balance font-bold text-[3em] leading-tight text-limelight">80%</h3>
                <p className="text-slate-300">Community Engagement</p>
              </div>
            </div>
          </Article>

          <div className="mt-40">Powered by @Kanicrafters</div>
        </main>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}

function Article({ children, className }: { children?: React.ReactNode; className?: string }) {
  return <article className={clsx('flex w-full mt-20 items-center justify-center', className)}>{children}</article>
}

function Title({ label, className }: { label: string; className?: string }) {
  return <h2 className={clsx('text-balance font-bold text-[4em] leading-tight', className)}>{label}</h2>
}
