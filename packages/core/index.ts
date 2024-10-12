import { makeInstaller } from '@zx-ui/utils'
import components from './components'

const installer = makeInstaller(components)

export * from '@zx-ui/components'

export default installer