import clsx from 'clsx'

export function Article({ children, className }: ArticleProps) {
  return <article className={clsx('flex w-full mt-20 items-center justify-center', className)}>{children}</article>
}

interface ArticleProps {
  children?: React.ReactNode
  className?: string
}
