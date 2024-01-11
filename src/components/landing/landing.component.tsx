import Image from 'next/image'
import { Article } from '@/components/article'
import { ButtonNeon } from '@/components/button'
import { Title } from '@/components/title'

import { TNewApplication } from '@/app/types'
import { createPool } from '@/sdk/allo'
import { createApplication } from '@/sdk/microgrants'
import { createProfile } from '@/sdk/registry'
import { Allocation } from '@allo-team/allo-v2-sdk/dist/strategies/MicroGrantsStrategy/types'
import { Status } from '@allo-team/allo-v2-sdk/dist/strategies/types'

export function Landing() {
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
    <Article>
      <div className="flex flex-col max-w-2xl">
        <Title label="Transparent, Secure, and Fueled by Innovation" />

        <p className="font-merienda text-2xl py-10">Decentralizing Dreams, Empowering Innovations.</p>
        <p className="text-slate-300 py-10">
          Explore, Fund, Decide - Revolutionizing collective financing with Arbitrum and Allo Protocol. Join us in a
          seamless journey of creation, funding, and decision-making.
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
  )
}
