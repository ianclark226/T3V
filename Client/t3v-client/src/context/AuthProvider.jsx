import React, { useEffect, useState, createContext } from 'react'

const AuthContext = createContext({})

export const AuthProvider = ({ children }) => {

    const [auth, setAuth] = useState(null)
    const [loading, setLoading] = useState(true)

    // Load auth from localStorage on mount
    useEffect(() => {
        const storedUser = localStorage.getItem('user')

        if (storedUser) {
            try {
                const parsedUser = JSON.parse(storedUser)
                setAuth(parsedUser)
            } catch (error) {
                console.error('Failed to parse user from local storage:', error)
                localStorage.removeItem('user')
            }
        }

        setLoading(false)   // <-- IMPORTANT: Always set loading false
    }, [])

    // Save auth to localStorage whenever it changes
    useEffect(() => {
        if (auth) {
            localStorage.setItem('user', JSON.stringify(auth))
        } else {
            localStorage.removeItem('user')
        }
    }, [auth])

    return (
        <AuthContext.Provider value={{ auth, setAuth, loading }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext