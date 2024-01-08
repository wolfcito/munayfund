import { TEAM } from '@/constants'
import { Article } from '@/components/article'
import { Title } from '@/components/title'
import Link from 'next/link'
import Image from 'next/image'

export function Team() {
  return (
    <Article className="flex-col">
      <Title label="Our team" className="mb-4" />

      <div className="flex gap-5">
        {TEAM.map(member => {
          return (
            <Link href={member.repo} key={crypto.randomUUID()}>
              <figure>
                <Image
                  src={member.image}
                  alt={`${member.name} avatar`}
                  width={100}
                  height={100}
                  className="h-60 w-auto rounded-lg"
                />
                <div>{member.name} </div>
              </figure>
            </Link>
          )
        })}
      </div>
    </Article>
  )
}
