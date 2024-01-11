import { ConnectButton } from '@rainbow-me/rainbowkit'
import Link from 'next/link'
import { ButtonNeon, ButtonSimple } from '../button'
import Image from 'next/image'

export function Header() {
  return (
    <div className="flex w-full items-center justify-between p-6">
      <Link href={'/'} className="text-limelight text-xl">
        <Image src={'/munayfund-logo.png'} alt="Munay logo" width={50} height={50} className="h-7 w-auto" />
      </Link>
      <div className='flex items-center'>
        <Link href={'/dashboard'} className="text-md flex mr-10 hover:underline">
          Dashboard
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
                    <div className="flex justify-center items-center">
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
                      <Link href={'/profile'} className="mx-3">
                        <div className="w-8 ring ring-limelight ring-offset-limelight ring-offset-1 rounded-lg">
                          <Image
                            src="/team/wolfcito.png"
                            alt="wolfcito profile"
                            width={33}
                            height={33}
                            className="rounded-lg"
                          />
                        </div>
                      </Link>
                    </div>
                  )
                })()}
              </div>
            )
          }}
        </ConnectButton.Custom>
      </div>
    </div>
  )
}
