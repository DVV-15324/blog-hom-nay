<<<<<<< HEAD

import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import { App } from "./App"

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        < App />
    </StrictMode>

)
=======
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
