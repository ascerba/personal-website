---
page: Blog
title: Personal PC Update
date: 2023-08-29
description: An update the new case mod for my personal desktop computer.
tags: blog computers
canonical: /blog
type: post
---

Playing with custom computers is a fun hobby. The common comparison is adult LEGO (as if adults don't play with LEGO...). To me, it's just another thing to be built, and I love a good thing to build. For a while now, I've had a Phanteks Enthoo Evolv ITX Tempered Glass Edition case, and while it does look decent, it is horribly space inefficient and hard to build in. Especially when you have a custom loop, which I do, accessing components on the motherboard is filled with frustration. It can be done, but you wonder why you're still in the hobby by the end. In any case, I acquired some rack-mount gear and a couple short wall-mount cabinets from a friend who recently moved out of town.

Rack mount computers have always been an interest of mine, but buying a case was always a bit out of the price range or unjustified because I already had a working case, if not a perfect one. With a rack, you can mount other stuff to the rails, too. And there are small 4U portable racks which make a nice all-in-one PC/Router/Access Point setup that you can take with you easier. A big thing for me recently is having a spot for everything, and more importantly, everything is grouped with items that are commonly used with eachother, and they're stored in something that is relatively portable. So having a 4U rack that has my PC and router for when I'm at school is something that appeals to be greatly. Also, rack mount PC's are much simpler to work on. The case is already on its side, and the components are all laid out in front of you with minimal overlap (I kinda messed this part up in my build, but hey, it's miles ahead of what I had before). This is honestly the biggest improvement over my Evolv ITX. Being able to swap out the RAM and CPU easily Is something I've wanted to be able to do for a while (it has come up multiple times recently). So, when I was able to get free rack equipment, I gladly accepted.

The interesting part is that the rack gear is old power supplies and scientific equipment, so it's not meant to house a computer. I've done some basic metal work before, so I figured it should be no problem to cut a few holes here and there. And so the adventure began.

![Initial mockup of the main components.](Mockup.jpg){alt="GPU, PSU, Motherboard, and radiator position mockup in the gutted power supply case."}

With the position of the components figured out, the first thing I wanted to get made and installed was the motherboard tray. With the tray installed, the exact position of the IO shield on the back plate could be figured out. Inside the old power supply was a bent piece of sheet aluminum that was used to hold some components together. Folded out, I figured it should have enough area for a motherboard.

![Cut layout for motherboard tray.](TrayLayout.jpg){alt="Marked cut lines on flattened aluminum."}

![Cuts made.](TrayCut.jpg){alt="Cut tray down to size."}

![Motherboard tray installed.](TrayInstalled.jpg){alt="Motherboard tray installed in case."}

With the tray cut and installed, the locations of the power supply and rear I/O could be determined and cut into the back plate.

![Laying out where the holes are](BackLayout.jpg){alt="Marking the layout of holes to cut on the rear panel."}

![Cut out I/O area.](IOShield.jpg){alt="Cut out location for and installed rear I/O shield."}

![Rear I/O shield fabricated.](IOandPSU.jpg){alt="Cut out PSU area and installed PSU."}

![PCIe slot cut and plate installed in case.](PlateInstalled.jpg){alt="Final completion of back plate with room for two-slot PCIe, I/O shield, and the PSU."}

The back plate was one of the more difficult parts of the build. Making sure cuts were accurate was very important in keeping things lined up. The last thing to do was to cut a hole up front for the air intake. By this point, however, I was more focused on packing for school, so I made the most jagged, terrible cut ever. But that's fine because it's still better than the almost non-existent airflow in the Evolv ITX. I also installed a dust filter up front, because I learned the hard way how much dust is in my dorm room last year, as well as a mesh plate to protect the two fans I plan to install outside the case eventually. The fans require a trip to the hardware store, and I didn't have time. Three fans will be plenty for now.

![Interior.](Top.jpg){alt="Inside of final build with custom loop liquid cooling."}

![Front.](Front.jpg){alt="Front of completed build."}

![Rear.](Rear.jpg){alt="Rear of completed build."}

![Top.](TopCover.jpg){alt="Top of case featuring original 'High Voltage' sticker from when it used to be a power supply."}

And with that, it was done. For the most part. When I got to school I decided that I should just install the second radiator. All I did was friction fit it in, but it's quite solid. The tubing is Tygon A-60-G (the stuff EK is trying to emulate with their ZMT line) and a metal anti-kink spring was used, because the 10/13mm tubing kinked very easily. The spring solved those issues, though, and I think it looks cooler with it.

![Final shot of inside with second radiator installed and system powered on.](Final.jpg){alt="Final interior shot with second radiator installed."}

Here are the specs:
- AMD Ryzen 9 5950X
- 32GB DDR4 3200MHz
- Nvidia Titan Xp (2017)
- Corsair RM850x
- Samsung 970 Evo Plus 512GB
- WD Blue 2TB 5400RPM
- EK Fluid Gaming custom loop kit
- 5x Corsair Light Loop 120mm

Some notes about the build:

Don't buy aluminum if you want a custom loop and aren't a machinist. There's not a good ecosystem of parts, and EK has pretty much abandoned their aluminum line. Also, I wouldn't buy Light Loops again. I just wanted them for a static color, but the RGB controller started freaking out after about a year of ownership, so I just removed it. Definitely get some "dumb" LED fans that don't need something like that, or just buy non-RGB fans. Also, I'm having instability with the 5950X. Lots of people seem to have a similar issue. The current fix is to make it run at base clock all the time, which is stupid. I may go back to my Ryzen 5 3600. Not sure if I can get AMD to RMA it since it was an eBay purchase from like a year ago, but I may try it. Basically, I wouldn't bother with the 5950X. Stick to the 3000 series.