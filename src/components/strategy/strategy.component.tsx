import { Article } from '@/components/article'
import { Title } from '@/components/title'
import Image from 'next/image'

export function Strategy() {
  return (
    <Article className="min-h-[600px]">
      <Image src="/munayfund.png" alt="logo" width={100} height={300} className="h-auto w-52 mr-20" />

      <div className="flex flex-col max-w-lg space-y-4">
        <Title label="Our very own Strategy" />

        <p className="text-xl text-slate-200">
          Introducing Cookie Jar Grants
        </p>
        <p className="text-slate-300">
          The cookie jar distribution strategy allows projects to recieve funding every x days once submitting a proof of delivery.
        </p>
        <br />
        <p className="text-slate-300">
          This prevents scammers from submitting projects, recieving funding and not actually building something.
          Proofs of delivery have to be approved by voters and only then can a project take their share of the cookie jar.
        </p>
      </div>

    </Article>
  )
}
