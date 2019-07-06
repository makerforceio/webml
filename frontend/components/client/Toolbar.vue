<template>
	<div class="bg-gray-200">
		<div class="
			head

			bg-primary text-white
			rounded-b-card
			p-2">
			<div class="m-4">
				<h1 class="
					font-bold
					text-4xl
				">Training title</h1>
				<div class="
					font-bold
				">By Author</div>
			</div>
			<div class="m-4"
				title="Major graph">
				<trend 
					:data="major"
					:gradient="gradient"
					:height="200"
					auto-draw
					smooth>
				</trend>
			</div>
			<div class="
				flex items-stretch
				m-4">
				<Button class="
					flex-1
					bg-white
					hover:bg-white
					text-primary">Pause</Button>
			</div>
		</div>
		<div class="p-2">
			<div class="
				bg-white
				rounded-card
				shadow-md
				m-4 p-2">
				<trend 
					:data="major"
					:height="200"
					auto-draw
					smooth>
				</trend>
			</div>
			<div class="
				bg-white
				rounded-card
				shadow-md
				m-4 p-2">
				<trend 
					:data="major"
					:height="200"
					auto-draw
					smooth>
				</trend>
			</div>
		</div>
	</div>
</template>

<style>
.head {
	background: linear-gradient(45deg, #FF6900, #FF994F);
	padding-bottom: 8rem;
	margin-bottom: -8rem;
}
</style>

<script>
import Button from '../common/Button.vue';
import trend from 'vuetrend';

const gradient = ['#ffffff', '#ff974d'];

export default {
	components: {
		Button,
		trend,
	},
	data() {
		return {
		gradient,
		major: [0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0],
		}
	}
};
</script>
