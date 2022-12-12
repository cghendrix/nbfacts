<script setup>

const { data } = await useFetch(() => `https://nbfacts.com/api/v1/facts/`)

function formatDateDay(date) {
  const options = {  weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' }
  return new Date(date).toLocaleDateString('en-us', options)
}

</script>

<template>
  <div>
    <header>
      <h1>Nickleback Facts</h1>
      <h5>Served via Google Cloud Platform - Text your fact to: </h5>
      <a href="tel:+12092082122">Text your favorite fact to 209.208.2122</a>
    </header>
    <main>
      <h2>Latest</h2>
      <ul v-if="data">
        <li v-for="fact in data" v-bind:key="fact.id">
          {{ fact.body }}
          <br/>
          <time>Added on {{ formatDateDay(fact.date_added) }}</time>
          <span id="via-sms" v-if="fact.sms_id !== null"> - Added via SMS</span>
        </li>
      </ul>
      <p v-else>No facts to show</p>
    </main>
  </div>
</template>
