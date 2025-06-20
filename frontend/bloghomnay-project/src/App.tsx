<<<<<<< HEAD
import { SnackbarProvider } from "notistack"
import { MainRoutes } from "./module/common/MainRoutes"
import { AuthProvider } from "./module/auth/context/authContext"
import { BrowserRouter } from "react-router-dom"
import "./index.css"



export const App = () => {
    return (

        <SnackbarProvider>
            <BrowserRouter>
                <AuthProvider>
                    <MainRoutes />
                </AuthProvider>
            </BrowserRouter>
        </SnackbarProvider>
    )
}

=======
import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
