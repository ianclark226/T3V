import { useLocation, Navigate, Outlet } from "react-router-dom";
import useAuth from "../hook/useAuth";
import { Spinner } from "react-bootstrap";

const RequiredAuth = () => {
    const { auth, loading } = useAuth()
    const location = useLocation()

    if(loading) {
        return (
            <Spinner/>
        )
    }

    return auth ? (
        <Outlet/>
    ) : (
        <Navigate to='/login' state={{from:location}} replace />
    )
}

export default RequiredAuth;