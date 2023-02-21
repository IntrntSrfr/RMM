<script setup lang="ts">
import useMessages from "../services/messages";

const { messages } = useMessages();
</script>

<template>
  <div>
    <p v-for="(m, i) in messages.data" :key="i">{{ m }}</p>
  </div>
</template>

<style lang="scss" scoped></style>
