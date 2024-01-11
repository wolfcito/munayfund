import style from './button.module.css'

export function ButtonNeon({
  onClick,
  label = 'Button',
}: {
  onClick: () => void
  label?: string
  children?: React.ReactNode
  className?: string
}) {
  return (
    <button className={style.button} onClick={onClick}>
      {label}
    </button>
  )
}
