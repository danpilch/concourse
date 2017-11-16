describe 'login', type: :feature do
  let(:team_name) { generate_team_name }
  let(:fly_home)  { Dir.mktmpdir }

  context 'to a team with space in its name' do
    it 'displays the team name correctly' do
      fly_login 'main'
      fly_with_input("set-team -n #{team_name} --no-really-i-dont-want-any-auth", 'y')
      fly_with_input("set-team -n \"#{team_name} test\" --no-really-i-dont-want-any-auth", 'y')

      visit dash_route("/teams/#{team_name}%20test/login")
      expect(page).to have_content "logging in to #{team_name} test"

      fly_with_input("destroy-team -n \"#{team_name} test\"", "#{team_name} test")
    end
  end
end
