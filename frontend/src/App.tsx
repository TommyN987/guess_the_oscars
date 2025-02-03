import { CategoryList } from "./components/CategoryList";
import background from "./assets/bbedb45f-14d1-4aee-a347-54a1fbd6e15c.webp";
import { useAuth } from "./hooks/useAuth";
import { Registration } from "./components/Registration";

function App() {
    const {
        context: { user },
    } = useAuth();

    return (
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
            <main className="w-full grow flex flex-col justify-center items-center">
                {user ? <CategoryList /> : <Registration />}
            </main>
        </div>
    );
}

export default App;
