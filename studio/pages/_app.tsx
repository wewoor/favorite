import '../styles/globals.css'
import type { AppProps } from 'next/app'
import '@vscode/codicons/dist/codicon.css';


function MyApp({ Component, pageProps }: AppProps) {
  return (
      <Component {...pageProps} />
  )
}

export default MyApp
