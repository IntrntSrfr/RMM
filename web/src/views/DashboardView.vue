<script setup lang="ts">
import TheDashboard from "@/components/TheDashboard.vue";
import TheNavbar from "@/components/TheNavbar.vue";
import AppScrollButton from "../components/AppScrollButton.vue";
</script>

<template>
  <TheNavbar />
  <Suspense>
    <TheDashboard />
  </Suspense>
  <AppScrollButton />
</template>
