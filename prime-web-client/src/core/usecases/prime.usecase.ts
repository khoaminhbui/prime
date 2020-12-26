import {PrimeRepository} from '../../infra/repos/prime.repo';

export class PrimeUsecase {
    async findLargestPrimeUnder(n: number) {
        const primeRepo = new PrimeRepository();
        return primeRepo.findLargestPrimeUnder(n);
    }
}