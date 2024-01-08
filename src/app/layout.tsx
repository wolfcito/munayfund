import './globals.css'

export default function RootLayout({ children }: LayoutProps) {
  return (
    <html lang="en">
      <body className={'font-poppins'}>{children}</body>
    </html>
  )
}

interface LayoutProps {
  readonly children: React.ReactNode
}
