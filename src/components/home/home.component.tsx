'use client'

import { TNewApplication } from '@/app/types'
import { createPool } from '@/sdk/allo'
import { createApplication } from '@/sdk/microgrants'
import { createProfile } from '@/sdk/registry'
import { chainData, wagmiConfigData } from '@/services/wagmi'
import { Allocation } from '@allo-team/allo-v2-sdk/dist/strategies/MicroGrantsStrategy/types'
import { Status } from '@allo-team/allo-v2-sdk/dist/strategies/types'
import { RainbowKitProvider, midnightTheme } from '@rainbow-me/rainbowkit'
import '@rainbow-me/rainbowkit/styles.css'
import { WagmiConfig } from 'wagmi'
import { ButtonNeon } from '@/components/button'
import Image from 'next/image'
import { Header } from '@/components/header'
import { Projects } from '@/components/projects'
import { Title } from '@/components/title'
import { Article } from '@/components/article/'
import Link from 'next/link'
import { TEAM } from '@/constants'

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
          <Article className="flex-col">
            <Title label="Projects Hub" />
            <Projects />
          </Article>

          <Article className="flex-col">
            <Title label="Our team" />

            <div className="flex gap-5">
              {TEAM.map(member => {
                return (
                  <Link href={member.repo} key={window.crypto.randomUUID()}>
                    <figure>
                      <Image
                        src={member.image}
                        alt={`${member.name} avatar`}
                        width={100}
                        height={100}
                        className="h-60 w-auto rounded-lg"
                      />
                      <div>{member.name} </div>
                    </figure>
                  </Link>
                )
              })}
            </div>
          </Article>

          <div className="self-center mt-40 mb-5">Powered by @Kanicrafters</div>
        </main>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}
