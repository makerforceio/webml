<template>
  <div>
    <div
      class="
        head

        bg-primary text-white
        rounded-b-card
        p-2"
    >
      <div class="max-w-4xl m-auto">
        <div class="flex items-center">
          <div class="flex-grow m-4">
            <h1
              class="
          font-bold
          text-4xl
        "
            >
              Training title
            </h1>
            <div
              class="
          font-bold
        "
            >
              By Author
            </div>
          </div>
          <Button class="m-4">
            <fa-icon size="2x" :icon="['far', 'pause']" />
          </Button>
        </div>
        <div class="mx-4 my-8" title="Major graph">
          <h2
            class="
			  font-medium"
          >
            Loss
          </h2>
          <trend
            :data="major"
            :gradient="gradient"
            :height="200"
            auto-draw
            smooth
          >
          </trend>
        </div>
      </div>
    </div>
    <div
      class="
  flex flex-wrap p-2
        max-w-4xl m-auto
    "
    >
      <div
        class="
    flex-grow
    flex
    flex-col
        bg-white
        rounded-card
        shadow-md
        m-4 py-4 px-6"
      >
        <h2
          class="
          font-medium"
        >
          Runtime
        </h2>
        <div
          class="
      flex-1
      flex
      justify-center
      items-center
          text-6xl font-bold
          text-center"
        >
          <span>10:55</span>
        </div>
      </div>
      <div
        class="
    flex-grow
    flex
    flex-col
        bg-white
        rounded-card
        shadow-md
        m-4 py-4 px-6"
      >
        <h2
          class="
          font-medium"
        >
          Runtime
        </h2>
        <div
          class="
      flex-1
      flex
      justify-center
      items-center
          text-6xl font-bold
          text-center"
        >
          <span>10:55</span>
        </div>
      </div>
      <div
        class="
    flex-grow
    flex
    flex-col
        bg-white
        rounded-card
        shadow-md
        m-4 py-4 px-6"
      >
        <h2
          class="
          font-medium"
        >
          Runtime
        </h2>
        <div
          class="
      flex-1
      flex
      justify-center
      items-center
          text-6xl font-bold
          text-center"
        >
          <span>10:55</span>
        </div>
      </div>
      <div
        class="
    flex-grow
        bg-white
        rounded-card
        shadow-md
        m-4 py-4 px-6"
      >
        <h2
          class="
          font-medium
          "
        >
          Epoch time
        </h2>
        <trend :data="major" :height="200" auto-draw smooth> </trend>
      </div>
      <div
        class="
    flex-grow
        bg-white
        rounded-card
        shadow-md
        m-4 py-4 px-6"
      >
        <h2
          class="
          font-medium
          "
        >
          Epoch time
        </h2>
        <trend :data="major" :height="200" auto-draw smooth> </trend>
      </div>
      <div
        class="
    flex-grow
        bg-white
        rounded-card
        shadow-md
        m-4 py-4 px-6"
      >
        <h2
          class="
          font-medium
          "
        >
          Epoch time
        </h2>
        <trend :data="major" :height="200" auto-draw smooth> </trend>
      </div>
    </div>
  </div>
</template>

<style>
.head {
  background: linear-gradient(45deg, #ff6900, #ff994f);
  padding-bottom: 6rem;
  margin-bottom: -6rem;
}
</style>

<script>
import Button from '~/components/common/Button.vue'
import trend from 'vuetrend'

const gradient = ['#ffffff', '#ff974d']

export default {
  layout: 'client',
  components: {
    Button,
    trend
  },
  data() {
    return {
      gradient,
      major: [0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0]
    }
  }
}
</script>
