'use client'
import { Header } from '@/components/header'
import { Team } from '@/components/team'
import { ProjectsHub } from '@/components/projects-hub'
import { MunayStats } from '@/components/munay-stats'
import { Landing } from '@/components/landing'

export function Home() {
  return (
    <main className="flex flex-col my-2 mx-7 min-h-screen">
      <Header />
      <Landing />
      <ProjectsHub />
      <MunayStats />
      <Team />
      <div className="self-center mt-40 mb-5">Powered by @Kanicrafters</div>
    </main>
  )
}
