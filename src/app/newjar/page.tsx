'use client'
import { Header } from '@/components/header'
import { createPool } from '@/sdk/allo'
import { useState } from 'react';

const NewJar = () => {
  const [poolId, setPoolId] = useState(null);

  return (
    <>
      <Header></Header>
      <h3 className="text-2xl font-semibold px-2 md:px-14">Create New Jar (Pool)</h3>
      <div className="flex items-center justify-center p-12">
        <div className="mx-auto w-full max-w-[700px]">
          <form action="" method="POST">
            <div className="-mx-3 flex flex-wrap">
              <div className="w-full px-3">
                <div className="mb-5">
                  <label htmlFor="endDate" className="mb-3 block text-base font-medium text-slate-200">
                    Submission Deadline
                  </label>
                  <input
                    type="date"
                    name="endDate"
                    id="endDate"
                    placeholder="2024-12-12"
                    className="w-full rounded-md border border-[#e0e0e0] bg-white py-3 px-6 text-base font-medium text-[#6B7280] outline-none focus:border-limelight focus:shadow-md"
                  />
                </div>
              </div>
              <div className="w-full px-3">
                <div className="mb-5">
                  <label htmlFor="votingDeadline" className="mb-3 block text-base font-medium text-slate-200">
                    Voting Deadline
                  </label>
                  <input
                    type="date"
                    name="votingDeadline"
                    id="votingDeadline"
                    placeholder="2024-12-31"
                    className="w-full rounded-md border border-[#e0e0e0] bg-white py-3 px-6 text-base font-medium text-[#6B7280] outline-none focus:border-limelight focus:shadow-md"
                  />
                </div>
              </div>
            </div>
            <div className="mb-5">
              <label htmlFor="distributionInterval" className="mb-3 block text-base font-medium text-slate-200">
                How many days should pass between payouts (interval)?
              </label>
              <input
                type="number"
                name="distributionInterval"
                id="distributionInterval"
                placeholder="15"
                min="15"
                step={15}
                className="w-full appearance-none rounded-md border border-[#e0e0e0] bg-white py-3 px-6 text-base font-medium text-[#6B7280] outline-none focus:border-limelight focus:shadow-md"
              />
            </div>

            <div className="mb-5">
              <label className="mb-3 block text-base font-medium text-slate-200">
                Do the projects require some data?
              </label>
              <div className="flex items-center space-x-6">
                <div className="flex items-center">
                  <input type="radio" name="radio1" id="radioButton1" className="h-5 w-5" />
                  <label htmlFor="radioButton1" className="pl-3 text-base font-medium text-slate-200">
                    Yes
                  </label>
                </div>
                <div className="flex items-center">
                  <input type="radio" name="radio1" id="radioButton2" className="h-5 w-5" />
                  <label htmlFor="radioButton2" className="pl-3 text-base font-medium text-slate-200">
                    No
                  </label>
                </div>
              </div>
            </div>

            <div>
              <button
                className="hover:shadow-form rounded-full bg-limelight py-3 px-8 text-center text-base font-semibold text-black outline-none"
                onClick={e => {
                  e.preventDefault()
                  createPool().then((res: any) => {
                    console.log('Pool ID: ', res.poolId)
                    alert('Pool created with ID: ' + res.poolId)
                    setPoolId(res.poolId)
                  })
                }}
              >
                Create
              </button>
            </div>
          </form>
        </div>
      </div>
    </>
  )
}

export default NewJar
