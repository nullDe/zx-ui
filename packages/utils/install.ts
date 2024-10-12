import type { App, Plugin} from "vue"

type SfCWithInstall<T> = T & Plugin

export function makeInstaller(comp: Plugin[]) {
  const installer = (app: App) => {
    comp.forEach(comp => app.use(comp))
  }

  return installer as Plugin
}

export function withInstall<T>(comp: T) {
  (comp as SfCWithInstall<T>).install = (app: App) => {
    const name = (comp as any).name

    app.component(name, comp as SfCWithInstall<T>)
  }

  return comp as SfCWithInstall<T>
}