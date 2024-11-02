// plugins/chapa.js
export default (context, inject) => {
    if (process.client) {
      const ChapaCheckout = require('@chapa_et/inline.js');
      inject('chapa', ChapaCheckout);
    }
  };
  