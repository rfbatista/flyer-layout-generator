import { create } from 'zustand'
import { Project } from './entities/project'

type Store = {
  projects: Project[]
}

const useProjectsStore = create<Store>(() => ({
  projects: [],
}))

export { useProjectsStore }
