import './globals.css'
import { chainData, wagmiConfigData } from '@/services/wagmi'
import { RainbowKitProvider } from '@rainbow-me/rainbowkit'
import { WagmiConfig } from 'wagmi'
import '@rainbow-me/rainbowkit/styles.css'

export default function RootLayout({ children }: LayoutProps) {
  return (
    <html lang="en">
      <body className={'font-poppins'}>
        <WagmiConfig config={wagmiConfigData}>
          <RainbowKitProvider chains={chainData} modalSize="compact" coolMode>
            {children}
          </RainbowKitProvider>
        </WagmiConfig>
      </body>
    </html>
  )
}

interface LayoutProps {
  readonly children: React.ReactNode
}
