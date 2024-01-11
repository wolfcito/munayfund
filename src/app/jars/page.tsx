'use client'
import { Header } from '@/components/header'
import React from 'react'
import Image from 'next/image'

const pools = [
  {
    name: 'Munay Genesis (First Round)',
    poolTotal: '0.5',
    payoutInterval: '30',
    image: '/projects/project001.jpg',
  },
  {
    name: 'TreeGrants',
    poolTotal: '100',
    payoutInterval: '15',
    image: '/projects/project002.jpg',
  },
]

type Props = {}

const Jars = (props: Props) => {
  return (
    <>
      <Header></Header>
      <div className="p-6 md:p-14">
        <div className="bg-[url('/projects/project001.jpg')] bg-no-repeat bg-cover rounded-2xl border border-limelight overflow-hidden mb-12 p-10 min-h-[400px]">
          <h3 className="text-2xl font-semibold pt-96">Newest</h3>
          <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Expedita iusto harum, eum cum deserunt impedit,
            sint minima natus ipsum corrupti fugit consequuntur possimus quam magni? Modi veniam autem consectetur
            voluptatibus.
          </p>
        </div>
        <div className="mb-10">
          <h3 className="text-2xl font-semibold">Active Jars</h3>
          <div className="mt-6 sm:grid sm:grid-cols-3 sm:gap-6">
            {pools.map((pool, index) => (
              <div
                key={index}
                className="rounded-2xl border border-limelight overflow-hidden col-span-2 md:col-span-1 mb-3"
              >
                <div className="h-[150px] bg-gradient-to-r from-limelight/55 via-lime-500 to-cyan-500 flex items-center justify-center">
                  <Image src={pool.image} alt="project image" width={60} height={60}></Image>
                </div>
                <div className="px-8 py-6">
                  <h4 className="text-base font-semibold leading-7 tracking-tight text-white">{pool.name}</h4>
                  <p className="text-sm leading-6 text-gray-400">Pool Total: {pool.poolTotal} ARB</p>
                  <p className="text-sm leading-6 text-gray-400">Payouts: each {pool.payoutInterval} days</p>
                </div>
              </div>
            ))}
          </div>
        </div>
        <div className="mb-10">
          <h3 className="text-2xl font-semibold">Past Jars</h3>
          <div className="mt-6 sm:grid sm:grid-cols-3 sm:gap-6">
            {pools.map((pool, index) => (
              <div
                key={index}
                className="rounded-2xl border border-limelight/15 overflow-hidden col-span-2 md:col-span-1 mb-3"
              >
                <div className="h-[150px] bg-gradient-to-r from-limelight/20 to-cyan-900 flex items-center justify-center">
                  <Image src={pool.image} alt="project image" width={60} height={60}></Image>
                </div>
                <div className="px-8 py-6">
                  <h4 className="text-base font-semibold leading-7 tracking-tight text-white">{pool.name}</h4>
                  <p className="text-sm leading-6 text-gray-400">Pool Total: {pool.poolTotal} ARB</p>
                  <p className="text-sm leading-6 text-gray-400">Payouts: each {pool.payoutInterval} days</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  )
}

export default Jars
