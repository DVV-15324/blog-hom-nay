import { SnackbarProvider } from "notistack"
import { MainRoutes } from "./module/common/MainRoutes"
import { AuthProvider } from "./module/auth/context/authContext"
import { BrowserRouter } from "react-router-dom"
import "./index.css"
import { TopLoadingBar } from "./module/common/TopLoadingBar"




export const App = () => {
  return (

    <SnackbarProvider>
      <BrowserRouter>
        <TopLoadingBar />
        <AuthProvider>
          <MainRoutes />
        </AuthProvider>
      </BrowserRouter>
    </SnackbarProvider>
  )
}

