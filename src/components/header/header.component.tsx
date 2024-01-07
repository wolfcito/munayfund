import { ConnectButton } from '@rainbow-me/rainbowkit'
import Link from 'next/link'
import { ButtonNeon, ButtonSimple } from '../button'
import Image from 'next/image'

export function Header() {
  return (
    <div className="flex w-full items-center justify-between">
      <Link href={'/'} className="text-limelight text-xl">
        MUNAYFUND
      </Link>
      <ConnectButton.Custom>
        {({ account, chain, openAccountModal, openChainModal, openConnectModal, authenticationStatus, mounted }) => {
          const ready = mounted && authenticationStatus !== 'loading'
          const connected =
            ready && account && chain && (!authenticationStatus || authenticationStatus === 'authenticated')

          return (
            <div
              {...(!ready && {
                'aria-hidden': true,
                style: {
                  opacity: 0,
                  pointerEvents: 'none',
                  userSelect: 'none',
                },
              })}
            >
              {(() => {
                if (!connected) {
                  return <ButtonNeon onClick={openConnectModal} label="Connect Wallet" />
                }

                if (chain.unsupported) {
                  return <ButtonNeon onClick={openChainModal} label=" Wrong network" />
                }

                return (
                  <div className="flex flex-col items-end">
                    <ButtonSimple onClick={openAccountModal} className="text-sm">
                      {account.displayName}
                      {account.displayBalance ? ` (${account.displayBalance})` : ''}
                    </ButtonSimple>
                    <ButtonSimple onClick={openChainModal} className="text-xs font-sans">
                      {chain.hasIcon && (
                        <div
                          style={{
                            background: chain.iconBackground,
                            width: 12,
                            height: 12,
                            borderRadius: 999,
                            overflow: 'hidden',
                            marginRight: 4,
                          }}
                        >
                          {chain.iconUrl ? (
                            <Image alt={chain.name ?? 'Chain icon'} src={chain.iconUrl} width={12} height={12} />
                          ) : null}
                        </div>
                      )}
                      {chain.name}
                    </ButtonSimple>
                  </div>
                )
              })()}
            </div>
          )
        }}
      </ConnectButton.Custom>
    </div>
  )
}
