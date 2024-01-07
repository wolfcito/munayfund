import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'MunayFund',
  description: 'The importance of support that goes beyond the material aspect',
}

export default function RootLayout({ children }: LayoutProps) {
  return (
    <html lang="en">
      <body className={'font-poppins'}>{children}</body>
    </html>
  )
}

interface LayoutProps {
  children: React.ReactNode
}
