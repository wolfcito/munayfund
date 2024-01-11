import { Article } from '@/components/article'
import { Projects } from '@/components/projects'
import { Title } from '@/components/title'

export function ProjectsHub() {
  return (
    <Article className="flex-col">
      <Title label="Projects Hub" />
      <Projects />
    </Article>
  )
}
