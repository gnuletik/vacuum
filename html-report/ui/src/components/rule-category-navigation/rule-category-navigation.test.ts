import { html } from 'lit';
import { fixture, expect } from '@open-wc/testing';
import { RuleCategoryNavigationComponent } from './rule-category-navigation-component';

describe('RuleCategoryNavigationComponent', () => {
  let element: RuleCategoryNavigationComponent;
  beforeEach(async () => {
    element = await fixture(
      html`<rule-category-navigation></rule-category-navigation>`
    );
  });

  it('renders nav', () => {
    const ul = element.shadowRoot!.querySelector('ul')!;
    expect(ul).to.exist;
  });

  it('passes the a11y audit', async () => {
    await expect(element).shadowDom.to.be.accessible();
  });
});
