package commands
/*
const Discord = require('discord.js');

const {embedColor} = require('../config.json');

const {getRandomInt,rollOutput} = require('../util.js');

module.exports = {
  name: 'regular',
  description: 'Roll a regular die',
  aliases: ['d'],
  execute(message, args) {
    var rollResults = []
    for (r = args.rolls; r > 0; r-- ) {
      rollResults.push(getRandomInt(1, args.sides));
    }

    result = rollOutput(resultList=rollResults, modifier=args.plusMinus)

    const resultEmbed = new Discord.MessageEmbed()
      .setColor(embedColor)

    if ( !result.total ) {
      resultEmbed
        .setTitle('Sorry!')
        .setDescription('Invalid command syntax')
    }
    else {
      resultEmbed
        .setTitle(result.total)
        .setFooter(result.formula || '')
    }

    message.reply(resultEmbed);
  }
};
*/

func execute() {
	var rollResults []int
	for i := 0
}