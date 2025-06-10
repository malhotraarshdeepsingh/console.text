import axios from 'axios';

// Pulsey API Endpoint URL To be added

export function Pulsey(key) {
    if (typeof key !== 'string') {
        throw new Error('API key must be a string')
    }

    return {
        async text(code, text) {
            if (typeof code !== 'string' || typeof text !== 'string') {
                throw new Error('Invalid input: code and text must be strings')
            }

            console.log(`Pulsey text: ${code}, ${text}`)

            try {
                const response = await axios.post('/pulsey', {
                    code,
                    text,
                    key
                });

                if (response.data?.error) {
                    return { success: false, error: response.data.error };
                }

                return response.data;
            } catch (error) {
                return { success: false, error: `Pulsey request failed: ${error.message}`}
            }
        },

        async siren(code, message) {
            if (typeof code !== 'string' || typeof message !== 'string') {
                throw new Error('Invalid input: code and message must be strings')
            }

            console.log(`Pulsey siren: ${code}, ${message}`);

            try {
                const response = await axios.post('/siren', {
                    code,
                    message,
                    key
                });

                if (response.data?.error) {
                    return { success: false, error: response.data.error };
                }

                return response.data;
            } catch (error) {
                return { success: false, error: `Pulsey siren request failed: ${error.message}` }
            }
        }
    }
}