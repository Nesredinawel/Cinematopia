const QRCode = require('qrcode');

async function generateQRCode(data) {
  try {
    const qrCodeUrl = await QRCode.toDataURL(data); // Generates a Data URL for the QR code image
    return qrCodeUrl;
  } catch (error) {
    console.error('Error generating QR code:', error);
    throw error; // Handle errors as needed
  }
}
