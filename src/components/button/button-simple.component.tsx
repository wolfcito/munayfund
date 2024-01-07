import clsx from 'clsx'

export function ButtonSimple({
  onClick,
  className,
  children,
}: {
  onClick: () => void
  label?: string
  children?: React.ReactNode
  className?: string
}) {
  return (
    <button className={clsx('flex items-center font-extralight hover:text-limelight', className)} onClick={onClick}>
      {children}
    </button>
  )
}
