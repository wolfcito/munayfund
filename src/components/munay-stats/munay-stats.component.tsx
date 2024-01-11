import Image from 'next/image'
import { Article } from '@/components/article'
import { Title } from '@/components/title'

export function MunayStats() {
  return (
    <Article className="min-h-[700px]">
      <div className="flex flex-col max-w-lg space-y-4">
        <Title label="About Munayfund" />

        <p className="text-slate-300">
          Our crowdfunding platform, powered by Arbitrum and Allo Protocol, revolutionizes decentralized funding,
          simplifying project creation and funding on the Arbitrum network.
        </p>
        <p className="text-slate-300">
          With a common fund managed by Allo Protocol, users vote for fund distribution, and the platform integrates
          machine learning, an intuitive interface, and robust security for transparent community engagement in
          decentralized finance evolution.
        </p>
      </div>

      <Image src="/upcurve.jpg" alt="levelup" width={100} height={300} className="h-auto w-52 ml-20" />
      <div className="px-4 space-y-5">
        <div className="text-right">
          <h3 className="text-balance font-bold text-[3em] leading-tight text-limelight">+2M</h3>
          <p className="text-slate-300">Funding volume</p>
        </div>
        <div className="text-right">
          <h3 className="text-balance font-bold text-[3em] leading-tight text-limelight">+50</h3>
          <p className="text-slate-300">Completed Projects</p>
        </div>
        <div className="text-right">
          <h3 className="text-balance font-bold text-[3em] leading-tight text-limelight">80%</h3>
          <p className="text-slate-300">Community Engagement</p>
        </div>
      </div>
    </Article>
  )
}
