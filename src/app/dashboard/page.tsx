'use client'
import { Header } from '@/components/header'
import Image from 'next/image'
import Link from 'next/link'

const pools = [
  {
    name: 'Munay Genesis (First Round)',
    poolTotal: '0.5',
    payoutInterval: '30',
    image: '/projects/project001.jpg'
  },
  {
    name: 'TreeGrants',
    poolTotal: '100',
    payoutInterval: '15',
    image: '/projects/project002.jpg'
  },
]

const Dashboard = () => {
  return (
    <div className="bg-black">
      <Header></Header>
      <div className="p-6 md:p-14">
        <div className="flex flex-row justify-between">
          <h3 className="text-2xl font-semibold">Created Pools</h3>
          <Link href="/newjar" className='border border-limelight rounded-full px-4 py-2 hover:shadow-md hover:shadow-lime-600'>New Pool</Link>
        </div>
        <div className='mt-6 sm:grid sm:grid-cols-3 sm:gap-6'>
          {pools.map((pool, index) => (
            <div key={index} className="rounded-2xl border border-limelight overflow-hidden col-span-2 md:col-span-1 mb-3">
              <div className='h-[150px] bg-gradient-to-r from-limelight/55 via-lime-500 to-cyan-500 flex items-center justify-center'>
                <Image src={pool.image} alt='project image' width={60} height={60}></Image>
              </div>
              <div className='px-8 py-6'>
                <h4 className="text-base font-semibold leading-7 tracking-tight text-white">{pool.name}</h4>
                <p className="text-sm leading-6 text-gray-400">Pool Total: {pool.poolTotal} ARB</p>
                <p className="text-sm leading-6 text-gray-400">Payouts: each {pool.payoutInterval} days</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default Dashboard
