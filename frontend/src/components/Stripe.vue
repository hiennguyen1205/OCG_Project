<template>
  <div>
    <div id="card-element">
      <input type="text" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Stripe',
  mounted() {
    const stripe = window.Stripe(
      'pk_test_51JFBJVAuOJBTYkNRKfuqR2SfLeoI7R1Ydkjc3yeBzESrHzR2fpS8uZi0i36lMRLNNaWN8GnTpBzP3bAgHxuFYNnG00Cgvss2Mp'
    );
    this.stripe = stripe;
    let elements = stripe.elements();
    this.card = elements.create('card', {
      hidePostalCode: true,
      style: {
        base: {
          iconColor: '#666EE8',
          color: '#31325F',
          lineHeight: '40px',
          fontWeight: 300,
          fontFamily: 'Helvetica Neue',
          fontSize: '15px',
          '::placeholder': {
            color: '#CFD7E0',
          },
        },
      },
    });
    this.card.mount('#card-element');
  },
  methods: {
    createPaymentMethod() {
      return this.stripe.createPaymentMethod('card', this.card);
    },
  },
};
</script>
