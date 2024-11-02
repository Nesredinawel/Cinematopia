const express = require('express');
const axios = require('axios');
const app = express();
const PORT = 8081;  // Or any port you prefer
const CHAPA_API_KEY = 'YOUR_CHAPA_API_KEY';

app.use(express.json());

// Proxy route for initializing a Chapa transaction
app.post('/api/initialize-chapa-transaction', async (req, res) => {
  try {
    const chapaResponse = await axios.post(
      'https://api.chapa.co/v1/transaction/initialize',
      req.body,
      {
        headers: {
          Authorization: `Bearer ${CHAPA_API_KEY}`,
          'Content-Type': 'application/json',
        },
      }
    );
    res.json(chapaResponse.data);  // Send Chapa response back to frontend
  } catch (error) {
    console.error('Error in Chapa initialization:', error);
    res.status(500).json({ error: 'Chapa initialization failed' });
  }
});

app.listen(PORT, () => console.log(`Chapa Proxy Server running on port ${PORT}`));
