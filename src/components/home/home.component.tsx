'use client'

import { chainData, wagmiConfigData } from '@/services/wagmi'
import { RainbowKitProvider, midnightTheme } from '@rainbow-me/rainbowkit'
import { WagmiConfig } from 'wagmi'
import { Header } from '@/components/header'
import { Team } from '@/components/team'
import { ProjectsHub } from '@/components/projects-hub'
import { MunayStats } from '@/components/munay-stats'
import { Landing } from '@/components/landing'
import '@rainbow-me/rainbowkit/styles.css'

export function Home() {
  return (
    <WagmiConfig config={wagmiConfigData}>
      <RainbowKitProvider chains={chainData} modalSize="compact" theme={midnightTheme()} coolMode>
        <main className="flex flex-col my-2 mx-7 min-h-screen">
          <Header />
          <Landing />
          <ProjectsHub />
          <MunayStats />
          <Team />
          <div className="self-center mt-40 mb-5">Powered by @Kanicrafters</div>
        </main>
      </RainbowKitProvider>
    </WagmiConfig>
  )
}
