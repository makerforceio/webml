<template>
  <div>
    <Header title="Sessions">
      <template v-slot:buttons-right>
        <button class="my-2 ml-4" @click="showNewSessionDialog = true">
          <fa-icon size="2x" :icon="['far', 'plus']" />
        </button>
      </template>
    </Header>
    <Cards>
      <Card
        v-for="model in models"
        :title="model.title"
        :key="model.title"
        arrow
      >
        <div class="flex">
          <Subcard subtitle="Elapsed">
            <CenteredText class="text-4xl">
              {{ model.elapsed }}
            </CenteredText>
          </Subcard>
          <Subcard subtitle="Loss">
            <CenteredText class="text-4xl">
              {{ model.loss }}
            </CenteredText>
          </Subcard>
        </div>
      </Card>
    </Cards>
    <NewSessionDialog :show.sync="showNewSessionDialog" />
  </div>
</template>

<script>
import Header from '~/components/common/Header.vue'
import Cards from '~/components/common/Cards.vue'
import Card from '~/components/common/Card.vue'
import Subcard from '~/components/common/Subcard.vue'
import CenteredText from '~/components/common/CenteredText.vue'
import NewSessionDialog from '~/components/admin/NewSessionDialog.vue'

export default {
  components: {
    Header,
    Cards,
    Card,
    Subcard,
    CenteredText,
    NewSessionDialog
  },
  data: () => ({
    showNewSessionDialog: false,
    models: [
      {
        title: 'Hello',
        elapsed: '10:43.4',
        loss: 43
      },
      {
        title: 'Hello',
        elapsed: '10:43.4',
        loss: 43
      }
    ]
  }),
  created: function() {
    // Initialize all the models and format them follow the data format above
	  const base = process.env.NUXT_ENV_BACKEND2_URL || 'http://localhost:10200';
    fetch(base + '/models').then((res) => {
      return res.json();
    }).then((body) => {
      return Promise.all(body.models
                        .filter(modelName => modelName != 'parser')
                        .map((modelName) => {
							return fetch(`${base}/params/${modelName}`)
                            .then((res) => res.text())
                            .then((loss) => { title: modelName, loss });
                        }));
    }).then((models) => {
      this.models = models;
    });
  }
}
</script>
