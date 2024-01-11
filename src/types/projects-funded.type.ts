export interface ProjectCollection {
  [key: string]: Project
}

interface Project {
  image: string
  'name-project': string
  'funded-percentage': string
  email: string
  'short-description': string
  description: string
}
