import React, {useState} from 'react';
import './styles.css';

function App() {
    const [inputValue, setInputValue] = useState('');
    const [jsonData, setJsonData] = useState(null);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(false);

    const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8085/v1/orders';

    const fetchData = async () => {
        if (!inputValue.trim()) return;

        setLoading(true);
        setError(null);
        setJsonData(null);

        try {
            const response = await fetch(`${API_BASE_URL}/${inputValue}`);
            const contentType = response.headers.get("content-type");
            let responseData = null;

            if (contentType && contentType.includes("application/json")) {
                responseData = await response.json();
            } else {
                responseData = await response.text();
            }

            if (!response.ok) {
                if (response.status === 404) {
                    setError(
                        typeof responseData === "object" && responseData !== null
                            ? responseData.message || "Заказ не найден"
                            : "Заказ не найден"
                    );
                } else {
                    setError(
                        typeof responseData === "object" && responseData !== null
                            ? responseData.message || `Ошибка ${response.status}`
                            : `Ошибка ${response.status}`
                    );
                }
                return;
            }

            setJsonData(responseData);
        } catch (err) {
            setError(err.message || "Произошла ошибка");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="container">
            <h1>Orders service</h1>
            <div className="input-section">
                <input
                    type="text"
                    value={inputValue}
                    onChange={(e) => setInputValue(e.target.value)}
                    placeholder="Введите ID"
                    disabled={loading}
                />
                <button onClick={fetchData} disabled={loading}>
                    {loading ? 'Загрузка...' : 'Получить данные'}
                </button>
            </div>

            {error && <div className="error">Ошибка: {error}</div>}

            {jsonData && (
                <div className="result">
                    <pre>{JSON.stringify(jsonData, null, 2)}</pre>
                </div>
            )}
        </div>
    );
}

export default App;
