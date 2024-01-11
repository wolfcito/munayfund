'use client'

import { useParams } from 'next/navigation'
import { Header } from '@/components/header'
import Image from 'next/image'
import { PROJECTS } from '@/constants'

const pools = [
  {
    name: 'Munay Genesis (First Round)',
    poolTotal: '0.5',
    payoutInterval: '30',
    image: '/pools/project001.jpg',
  },
  {
    name: 'TreeGrants',
    poolTotal: '100',
    payoutInterval: '15',
    image: '/pools/project002.jpg',
  },
]

export default function JarDetails() {
  const params = useParams<{ id: string }>()
  const { id } = params
  const intId = parseInt(id) - 1

  return (
    <>
      <Header></Header>
      <div className="bg-black">
        <div className={`bg-[url('/pools/pool001.jpg')] bg-no-repeat bg-cover h-[50vh]`}>
          <div className="h-[50vh] bg-gradient-to-b from-black/0 to-black/100"></div>
        </div>
        <div className="p-6 md:p-14 mt-[-100px] max-w-[800px]">
          <div>
            <h3 className="text-4xl font-semibold">{pools[intId].name}</h3>
            <p className="mt-1  leading-6 text-gray-300">Total: {pools[intId].poolTotal} ARB</p>
          </div>
          <div>
            <dl>
              <div className="py-6 sm:grid sm:grid-cols-3 sm:gap-4">
                <dt className=" font-medium">DaysInterval</dt>
                <dd className="mt-1  leading-6 text-gray-300 sm:col-span-2 sm:mt-0">{pools[intId].payoutInterval}</dd>
              </div>
              
              <h3 className="text-2xl font-semibold">Projects in this Jar</h3>
              <div className="mt-6 sm:grid sm:grid-cols-3 sm:gap-6">

                  <div
                    className="rounded-2xl border border-limelight overflow-hidden col-span-2 md:col-span-1 mb-3"
                  >
                    <div className="h-[150px] bg-gradient-to-r from-limelight/55 via-lime-500 to-cyan-500 flex items-center justify-center">
                      <Image src={PROJECTS[1]['image']} alt="project image" width={60} height={60}></Image>
                    </div>
                    <div className="px-8 py-6">
                      <h4 className="text-base font-semibold leading-7 tracking-tight text-white">{PROJECTS[1]['name-project']}</h4>
                      <p className="text-sm leading-6 text-gray-400">Funded: {PROJECTS[1]['funded-percentage']}</p>
                    </div>
                  </div>

              </div>
            </dl>
          </div>
        </div>
      </div>
    </>
  )
}
