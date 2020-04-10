<template>
  <div>
    
    
    <div @click="drawer = true">
      <Animition />
    </div>
    <h2 style="text-align:right;margin:40px;">
      当前版本号：v0.1.0
    </h2>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Animition from '@/view/dashbord/component/animition.vue'
export default {
  name: 'Dashbord',
  data() {
    return {
      drawer: false
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  components: {
    Animition
  }
}
</script>

<style lang="scss" scoped>
</style>