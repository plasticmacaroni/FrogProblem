# FrogProblem
This is a demonstration of the frog problem, which you can check out on YouTube [here](https://www.youtube.com/watch?v=cpwSGsb-rTs).

--

The basic premise of the problem is this:
You've eaten a deadly mushroom, and the only cure is to lick a female frog. Only male frogs croak, so if you hear a croak, you know that a male (which can't cure your poisoning) is present. You're lost in the rainforest, and don't have enough time to lick frogs in different places, but you spy two areas with frogs:
* One singular frog, sitting in front of you
* Two frogs together, behind you (however, you hear a croak from one of these frogs, and aren't sure which frog croaked)

If you go for the lone frog, you have a 1/2 chance of it being female. However, if you go for the two frogs together, the video posits you have a 2/3 chance of the frog being female. This seems counterintuitive, but I attempted to show that this is a valid statistical figure.

--

This program creates groups of two frogs, and then calculates the odds of one of the frogs being female.

This shows that 66% of the time (2/3), there is a female in the group. 

It's important to note that we don't know *which* of the frogs croaked. If we did know which frog croaked, the odds would be closer to 50%.

I didn't believe this myself, hence writing the program to hopefully demonstrate things objectively.
