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
    return (
        <Router>
            <div
                style={{
                    backgroundImage: `url(${background})`,
                    backgroundSize: "cover",
                    backgroundPosition: "center",
                }}
                className="min-h-screen flex flex-col items-center bg-gradient-to-b from-black via-gray-900 to-black text-white font-sans"
            >
                <div className="absolute inset-0 bg-black opacity-70" />
                <header className="w-full flex justify-center font-serif bg-black text-4xl opacity-80 z-20 p-2 text-gold">
                    <h1>Guess the Oscars</h1>
                </header>
                <Routes>
                    <Route path="/" element={<Main />} />
                    <Route path="/validate" element={<EmailValidation />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
