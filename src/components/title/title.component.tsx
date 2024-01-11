import clsx from 'clsx'

export function Title({ label, className }: TitleProps) {
  return <h2 className={clsx('text-balance font-bold text-[4em] leading-tight', className)}>{label}</h2>
}

interface TitleProps {
  label: string
  className?: string
}
