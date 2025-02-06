import { CategoryList } from "./components/CategoryList";
import background from "./assets/bbedb45f-14d1-4aee-a347-54a1fbd6e15c.webp";
import { useAuth } from "./hooks/useAuth";
import { Registration } from "./components/Registration";
import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import { EmailValidation } from "./components/EmailValidation";

function Main() {
    const {
        context: { user, loading },
    } = useAuth();

    if (loading) {
        return <div className="text-white text-xl">Loading...</div>; // Show loading state
    }

    return (
        <main className="w-full grow flex flex-col justify-center items-center">
            {user ? <CategoryList /> : <Registration />}
        </main>
    );
}

function App() {
    const {
        context: { user, setUser },
        logout,
    } = useAuth();

    const handleLogout = async () => {
        await logout();
        setUser(null);
    };

    return (
        <Router>
            <div
                style={{
                    backgroundImage: `url(${background})`,
                    backgroundSize: "cover",
                    backgroundPosition: "center",
                }}
                className="fixed inset-0 min-h-screen flex flex-col items-center bg-gradient-to-b from-black via-gray-900 to-black text-white font-sans"
            >
                <div className="absolute inset-0 bg-black opacity-60" />
                <header className="w-full flex justify-center font-serif bg-black text-2xl sm:text-4xl opacity-90 p-2 text-gold">
                    <h1 className="z-40">Guess the Oscars</h1>
                </header>
                {user && (
                    <p className="sm:block absolute top-3 left-5 px-4 py-1 font-bold text-gold z-20 hidden">{`Welcome, ${user.name}`}</p>
                )}
                <button
                    className="hidden sm:block absolute top-3 right-5 px-4 py-1 bg-gold text-black border-2 border-gold rounded-lg cursor-pointer hover:text-gold hover:bg-black z-20"
                    onClick={handleLogout}
                >
                    Logout
                </button>
                <div className="relative w-screen min-h-screen flex flex-col items-center text-white font-sans overflow-auto z-10">
                    <Routes>
                        <Route path="/" element={<Main />} />
                        <Route path="/validate" element={<EmailValidation />} />
                    </Routes>
                </div>
            </div>
        </Router>
    );
}

export default App;
