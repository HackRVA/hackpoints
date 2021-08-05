# Hack Points
Hack Points is a project created for hackrva (hackerspace).  

## Run it
We are using a Makefile.

You need [golang](https://golang.org/doc/install) installed (preferably v1.16).

```bash
make run
```
> this will start the app on localhost:3000 (however, there's no frontend yet)

## Run Tests
```bash
make test
```

## Bounties
A bounty is a task or action item that needs to get done.

- A member can create a bounty
- Member adds some details so that everyone understands what the bounty is
- Members will have some time to assign a point value to each bounty -- See [points](#points)
- The bounty can be completed by one person or multiple people.
- When the bounty is done, the space score goes up.
- at a certain score, we have a celebratory pizza party

## Points
We determine how many points a bounty is worth as a group.
After a bounty is created, it will enter a period of time where people can endorse it.
Each endorsement is worth one point.

Endorse wisely.
> might be worth considering limiting how many endorsements per month a member can dish out.

## Design Considerations
checkout [design considerations](./docs/design_considerations.md)!
