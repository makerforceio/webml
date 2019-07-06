<template>
	<div class="
		bg-primary text-white
		rounded-b-card"
		style="background: linear-gradient(30deg, #FF6900, #FF974D)">
		<div class="px-4 py-8">
			<h1 class="
				font-bold
				text-xl
			">Training title</h1>
		</div>
		<div class="p-4"
			title="Major graph">
			<trend 
				:data="major"
				:gradient="gradient"
				height="200"
				auto-draw
				smooth>
			</trend>
		</div>
		<div class="
			flex items-stretch
			p-4">
			<Button class="
				flex-1
				bg-white
				hover:bg-white
				text-primary">Pause</Button>
		</div>
	</div>
</template>

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
