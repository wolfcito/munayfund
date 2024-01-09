'use client'

import { useParams } from 'next/navigation'
import { PaperClipIcon } from '@heroicons/react/20/solid'
import { PROJECTS } from '@/constants'

export default function ProjectProfile() {
  const params = useParams<{ id: string }>()
  const { id } = params

  return (
    <div className="bg-[url('/projects/project001.jpg')] bg-no-repeat bg-cover">
      <div className="bg-black/85">
        <div className="px-4 py-6 sm:px-6 ">
          <h3 className="text-xl font-semibold leading-7">Applicant Information</h3>
          <p className="mt-1 max-w-2xl text-sm leading-6 text-gray-300">Project details</p>
        </div>

        <dl>
          <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium">Name project</dt>
            <dd className="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0">
              {PROJECTS[id]['name-project']}
            </dd>
          </div>
          <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium">Funded %</dt>
            <dd className="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0">
              {PROJECTS[id]['funded-percentage']}
            </dd>
          </div>
          <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium">Email address</dt>
            <dd className="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0">{PROJECTS[id]['email']}</dd>
          </div>
          <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium">Short description</dt>
            <dd className="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0">
              {PROJECTS[id]['short-description']}
            </dd>
          </div>
          <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium">Description</dt>
            <dd className="mt-1 text-sm leading-6 text-gray-300 sm:col-span-2 sm:mt-0">
              {PROJECTS[id]['description']}
            </dd>
          </div>
          <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium leading-6">Attachments</dt>
            <dd className="mt-2 text-sm sm:col-span-2 sm:mt-0">
              <ul role="list" className="divide-y divide-gray-700 rounded-md border border-gray-200">
                <li className="flex items-center justify-between py-4 pl-4 pr-5 text-sm leading-6">
                  <div className="flex w-0 flex-1 items-center">
                    <PaperClipIcon className="h-5 w-5 flex-shrink-0 text-gray-300" aria-hidden="true" />
                    <div className="ml-4 flex min-w-0 flex-1 gap-2">
                      <span className="truncate font-medium">Interview_Audio_Log.wav.pdf</span>
                      <span className="flex-shrink-0 text-gray-300">2.4mb</span>
                    </div>
                  </div>
                  <div className="ml-4 flex-shrink-0">
                    <a href="#" className="font-medium text-limelight">
                      Download
                    </a>
                  </div>
                </li>
                <li className="flex items-center justify-between py-4 pl-4 pr-5 text-sm leading-6">
                  <div className="flex w-0 flex-1 items-center">
                    <PaperClipIcon className="h-5 w-5 flex-shrink-0 text-gray-300" aria-hidden="true" />
                    <div className="ml-4 flex min-w-0 flex-1 gap-2">
                      <span className="truncate font-medium">Forensic_Analysis_Report.pdf</span>
                      <span className="flex-shrink-0 text-gray-300">4.5mb</span>
                    </div>
                  </div>
                  <div className="ml-4 flex-shrink-0">
                    <a href="#" className="font-medium text-limelight">
                      Download
                    </a>
                  </div>
                </li>
              </ul>
            </dd>
          </div>
        </dl>
      </div>
    </div>
  )
}
